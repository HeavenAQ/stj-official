import React from "react";
import emailjs from "@emailjs/browser";
import toast, { Toaster } from "react-hot-toast";
import "react-phone-number-input/style.css";
import PhoneInput from "react-phone-number-input";
import SakeOrderCards from "./SakeOrderCards";
import { sakes } from "../data/sakes";
import { Order } from "../types";
import { usePlacesWidget } from "react-google-autocomplete";
import Loading from "./Loading";

const serviceID = process.env.REACT_APP_EMAILJS_SERVICE_ID;
const templateID = process.env.REACT_APP_EMAILJS_TEMPLATE_ID;
const apiKey = process.env.REACT_APP_EMAILJS_API_KEY;
const googleMapAPIKey = process.env.REACT_APP_GOOGLE_MAP_API_KEY;

const initOrders = (sakes: string[]) => {
  return sakes.map((sake) => {
    return {
      name: sake,
      quantity: 0,
      note: "",
    };
  });
};

const ContactUs: React.FC = () => {
  const [name, setName] = React.useState<string>("");
  const [email, setEmail] = React.useState<string>("");
  const [phone, setPhone] = React.useState<any>();
  const [address, setAddress] = React.useState<string>("");
  const [backupAddress, setBackupAddress] = React.useState<string>("");
  const [spinner, setSpinner] = React.useState<boolean>(false);

  // initialize the order list
  const [orders, setOrders] = React.useState<Order[]>(initOrders(sakes.chn));

  // setup google map API
  const { ref } = usePlacesWidget<HTMLInputElement>({
    apiKey: googleMapAPIKey,
    onPlaceSelected: (place) => {
      setAddress(place.formatted_address);
    },
    options: {
      types: ["(regions)"],
      componentRestrictions: { country: "tw" },
    },
  });

  // send email to the admin when form submitted
  const onSubmit = (ev: React.FormEvent) => {
    ev.preventDefault();
    setSpinner(true);
    const finalAddress = address || backupAddress;

    // check if all fields are filled
    if (
      !(
        name &&
        email &&
        phone &&
        finalAddress &&
        orders.some((order) => order.quantity > 0)
      )
    ) {
      setSpinner(false);
      toast.error("請輸入所有項目，並且確認物品數量不為0，謝謝！");
      return;
    }

    // set email template parameters
    const templateParams = {
      name: name,
      from_email: email,
      phone: phone,
      address: finalAddress,
      body: orders
        .filter((order) => order.quantity > 0)
        .map((order) => {
          return `${order.name} x ${order.quantity} *備註：${order.note}`;
        })
        .join("\n"),
    };

    // ensure the emailjs service ID, template ID, and API key are set
    if (!serviceID || !templateID || !apiKey) {
      toast.error("送出失敗，請再試一次！");
      console.log("EmailJS service ID, template ID, or API key is not set.");
      setSpinner(false);
      return;
    }

    // send email
    emailjs
      .send(serviceID, templateID, templateParams, apiKey)
      .then((_) => {
        setSpinner(false);
        setName("");
        setEmail("");
        setPhone("");
        setOrders(initOrders(sakes.chn));
        toast.success("已成功送出！謝謝您的訂購！");
      })
      .catch((res) => {
        setSpinner(false);
        console.log("Failed to send email.", res);
        toast.error("送出失敗，請再試一次！");
      });
  };

  return (
    <section className="pt-8 w-full h-auto bg-gray-100 rounded-xl">
      <div className="pt-10 mx-auto w-full md:w-4/5 lg:w-2/3 max-w-[1000px] h-[950px] md:h-[850px]">
        <h1 className="mb-8 text-2xl font-black text-center lg:mb-12 lg:text-3xl tracking-[20px]">
          聯絡我們
        </h1>
        {spinner && (
          <div className="flex justify-center items-center mx-auto mt-20 w-32 h-32">
            <Loading />
          </div>
        )}
        <form
          className={`${spinner ? "hidden" : ""}  p-6 w-full rounded-xl`}
          onSubmit={onSubmit}
        >
          <div className="flex flex-wrap -mx-3 mb-6">
            <div className="px-3 mb-6 w-full md:mb-0 md:w-1/2">
              <input
                className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
                id="grid-first-name"
                type="text"
                placeholder="姓名"
                onChange={(ev) => setName(ev.target.value)}
                value={name}
                required
              />
            </div>
            <div className="px-3 w-full md:w-1/2">
              <input
                className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
                id="email"
                type="email"
                placeholder="example@mail.com"
                onChange={(ev) => setEmail(ev.target.value)}
                value={email}
                required
              />
            </div>
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-6">
            <PhoneInput
              className="block py-3 px-4 mb-3 w-full leading-tight bg-white rounded border appearance-none"
              placeholder="Enter phone number"
              defaultCountry="TW"
              value={phone}
              onChange={setPhone}
            />
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-6">
            <input
              ref={ref}
              className="block py-3 px-4 mb-3 w-full leading-tight bg-white rounded border appearance-none"
              placeholder="請輸入地址"
              type="text"
              onChange={(ev) => setBackupAddress(ev.currentTarget.value)}
            />
          </div>
          <div className="flex flex-wrap -mx-3 mb-10">
            <SakeOrderCards orders={orders} setOrders={setOrders} />
          </div>

          <div className="flex flex-wrap -mx-3">
            <div className="flex justify-center items-end w-full h-auto text-white">
              <button
                type="submit"
                className="inline-flex overflow-hidden relative justify-start items-center py-3 px-6 font-medium rounded-xl transition-all bg-zinc-500 group"
              >
                <span className="inline-block absolute top-0 right-0 w-4 h-4 rounded transition-all duration-500 ease-in-out group-hover:-mt-4 group-hover:-mr-4 bg-zinc-700">
                  <span className="absolute top-0 right-0 w-5 h-5 bg-white rotate-45 translate-x-1/2 -translate-y-1/2"></span>
                </span>
                <span className="absolute bottom-0 left-0 w-full h-full rounded-2xl transition-all duration-500 ease-in-out delay-200 -translate-x-full translate-y-full group-hover:mb-12 group-hover:translate-x-0 bg-zinc-600"></span>
                <span className="relative w-full text-left text-white transition-colors duration-200 ease-in-out group-hover:text-white">
                  送出
                </span>
              </button>
            </div>
          </div>
        </form>
      </div>
    </section>
  );
};

export default ContactUs;

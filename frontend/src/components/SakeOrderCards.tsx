import React, { useRef } from "react";
import { MdArrowBackIos, MdArrowForwardIos } from "react-icons/md";
import { sakes } from "../data/sakes";
import { Order } from "../types";

const SakeOrderCard = (
  imageURL: string,
  key: number,
  orders: Order[],
  setOrders: React.Dispatch<React.SetStateAction<Order[]>>,
) => {
  const updateQuantity = (key: number, quantity: number) => {
    const newOrders = [...orders];
    newOrders[key].quantity = quantity;
    setOrders(newOrders);
  };

  return (
    <div key={key} className="flex flex-col items-center">
      <img
        src={imageURL}
        alt="sake"
        className="object-cover w-60 h-60 rounded-lg"
      />
      <div className="flex flex-col items-center mt-5 space-y-3 w-44">
        <input
          type="text"
          className="py-2 px-4 mt-2 w-full bg-white rounded-lg"
          placeholder="備註："
          value={orders[key].note}
          onChange={(ev) => {
            const newOrders = [...orders];
            newOrders[key].note = ev.target.value;
            setOrders(newOrders);
          }}
        />
        <div className="inline-flex justify-evenly items-center w-full">
          <button
            className="w-8 h-8 text-white rounded-lg bg-zinc-600"
            type="button"
            onClick={() => updateQuantity(key, orders[key].quantity - 1)}
          >
            -
          </button>
          {orders[key].quantity}
          <button
            className="w-8 h-8 text-white rounded-lg bg-zinc-600"
            type="button"
            onClick={() => updateQuantity(key, orders[key].quantity + 1)}
          >
            +
          </button>
        </div>
      </div>
    </div>
  );
};

interface SakeOrderCardsProps {
  orders: Order[];
  setOrders: React.Dispatch<React.SetStateAction<Order[]>>;
}

const SakeOrderCards: React.FC<SakeOrderCardsProps> = ({
  orders,
  setOrders,
}) => {
  const sakeInfos = sakes.getSakeInfoList();
  const yearRef = useRef<HTMLDivElement>(null);
  const onClick = (direction: "left" | "right") => {
    // ensure the container exists
    const container = yearRef.current as HTMLDivElement;

    // calculate the new scroll position
    const containerWidth = container.clientWidth;
    const newScrollPosition =
      direction === "right"
        ? container.scrollLeft + containerWidth
        : container.scrollLeft - containerWidth;

    // scroll to the new position
    container.scrollTo({
      top: 0,
      left: newScrollPosition,
      behavior: "smooth",
    });
  };

  return (
    <div
      className="flex overflow-y-hidden overflow-x-scroll flex-nowrap mx-auto snap-mandatory snap-x max-w-[1000px] scroll-smooth"
      ref={yearRef}
    >
      {sakeInfos.map((sake, i) => {
        return (
          <div className="flex-[0_0_100%] snap-start sm:px-10 relative" key={i}>
            <MdArrowBackIos
              className="absolute left-1.5 top-[78%] text-xl cursor-pointer md:top-1/2 md:-translate-y-1/2 text-zinc-600"
              onClick={() => onClick("left")}
            />
            <MdArrowForwardIos
              className="absolute right-1 top-[78%] text-xl cursor-pointer md:top-1/2 md:-translate-y-1/2 text-zinc-600"
              onClick={() => onClick("right")}
            />
            <div className="h-full">
              {SakeOrderCard(sake.images[0], i, orders, setOrders)}
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default SakeOrderCards;

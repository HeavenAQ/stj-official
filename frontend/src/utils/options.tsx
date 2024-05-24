export const YearOptions = () => {
  const currentYear = new Date().getFullYear()
  const years = Array.from({ length: 82 }, (_, i) => currentYear - i - 18)
  return (
    <>
      {years.map(year => (
        <option key={year} value={year}>
          {year}
        </option>
      ))}
    </>
  )
}

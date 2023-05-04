import React from 'react'

interface AlgorithmChooserProps {
  isKMP: boolean
  setIsKMP: (newVal: boolean) => void
}

const AlgorithmChooser: React.FC<AlgorithmChooserProps> = ({
  isKMP, setIsKMP
}) => {
  const handleRadioChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
    setIsKMP(event.target.value === 'KMP')
  }

  return (
    <div className="bg-secondary-base w-60 px-4">
      <label className="label justify-start font-bold">Choose Algorithm</label>
      <div className="form-control">
        <label className="label cursor-pointer justify-start">
          <input
            type="radio"
            name="radio-10"
            className="radio checked:bg-primary-base"
            value="KMP"
            checked={isKMP}
            onChange={handleRadioChange}
          />
          <label className="label-text ml-2 text-white">Knuth-Morris-Pratt (KMP)</label>
        </label>
      </div>
      <div className="form-control">
        <label className="label cursor-pointer justify-start">
          <input
            type="radio"
            name="radio-10"
            className="radio checked:bg-primary-base"
            value="BM"
            checked={!isKMP}
            onChange={handleRadioChange}
          />
          <label className="label-text ml-2 text-white">Boyer-Moore (BM)</label>
        </label>
      </div>
    </div>
  )
}

export default AlgorithmChooser

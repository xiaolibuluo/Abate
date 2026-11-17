using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public static int Scale13020<T>(T value) where T : struct
    {
        int number = 53;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
    public static int Scale13021<T>(T value) where T : struct
    {
        int number = 60;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 45; }
        return number;
    }
    public struct Point13022
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 67;
    }

    public static class Grid13022
    {
        public static Point13022 Make(int left, int right)
        {
            return new Point13022 { X = left, Y = right };
        }
    }
    public delegate void Notifier13023(int value);

    public sealed class Source13023
    {
        public event Notifier13023 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale13024<T>(T value) where T : struct
    {
        int number = 81;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
    public sealed class SlotMap13025
    {
        private readonly Dictionary<int, string> items = new Dictionary<int, string>();

        public string this[int index]
        {
            get => items.TryGetValue(index, out var value) ? value : string.Empty;
            set => items[index] = value;
        }

        public string Describe(int index) => index switch
        {
            0 => "zero",
            _ => items.ContainsKey(index) ? "present" : "absent"
        };
    }
}

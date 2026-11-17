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

    public struct Point21360
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 31;
    }

    public static class Grid21360
    {
        public static Point21360 Make(int left, int right)
        {
            return new Point21360 { X = left, Y = right };
        }
    }
    public sealed class SlotMap21361
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
    public static int Scale21362<T>(T value) where T : struct
    {
        int number = 45;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 38; }
        return number;
    }
    public static int Scale21363<T>(T value) where T : struct
    {
        int number = 52;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public static int Scale21364<T>(T value) where T : struct
    {
        int number = 59;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 60; }
        return number;
    }
    public delegate void Notifier21365(int value);

    public sealed class Source21365
    {
        public event Notifier21365 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

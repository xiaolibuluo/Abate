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

    public sealed class SlotMap17800
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
    public delegate void Notifier17801(int value);

    public sealed class Source17801
    {
        public event Notifier17801 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale17802<T>(T value) where T : struct
    {
        int number = 73;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public struct Point17803
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 80;
    }

    public static class Grid17803
    {
        public static Point17803 Make(int left, int right)
        {
            return new Point17803 { X = left, Y = right };
        }
    }
}

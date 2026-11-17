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

    public sealed class SlotMap21810
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
    public delegate void Notifier21811(int value);

    public sealed class Source21811
    {
        public event Notifier21811 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap21812
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
    public static int Scale21813<T>(T value) where T : struct
    {
        int number = 55;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 59; }
        return number;
    }
    public struct Point21814
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 62;
    }

    public static class Grid21814
    {
        public static Point21814 Make(int left, int right)
        {
            return new Point21814 { X = left, Y = right };
        }
    }
}

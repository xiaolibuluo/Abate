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

    public static int Scale24220<T>(T value) where T : struct
    {
        int number = 63;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 66; }
        return number;
    }
    public delegate void Notifier24221(int value);

    public sealed class Source24221
    {
        public event Notifier24221 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point24222
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 77;
    }

    public static class Grid24222
    {
        public static Point24222 Make(int left, int right)
        {
            return new Point24222 { X = left, Y = right };
        }
    }
    public sealed class SlotMap24223
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
    public delegate void Notifier24224(int value);

    public sealed class Source24224
    {
        public event Notifier24224 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

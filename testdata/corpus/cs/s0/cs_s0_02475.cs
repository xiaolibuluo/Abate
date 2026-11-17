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

    public struct Point23110
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 75;
    }

    public static class Grid23110
    {
        public static Point23110 Make(int left, int right)
        {
            return new Point23110 { X = left, Y = right };
        }
    }
    public delegate void Notifier23111(int value);

    public sealed class Source23111
    {
        public event Notifier23111 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale23112<T>(T value) where T : struct
    {
        int number = 89;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 43; }
        return number;
    }
    public sealed class SlotMap23113
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

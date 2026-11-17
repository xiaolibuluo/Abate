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

    public struct Point4820
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 63;
    }

    public static class Grid4820
    {
        public static Point4820 Make(int left, int right)
        {
            return new Point4820 { X = left, Y = right };
        }
    }
    public static int Scale4821<T>(T value) where T : struct
    {
        int number = 70;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
    public sealed class SlotMap4822
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
    public delegate void Notifier4823(int value);

    public sealed class Source4823
    {
        public event Notifier4823 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point4824
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 91;
    }

    public static class Grid4824
    {
        public static Point4824 Make(int left, int right)
        {
            return new Point4824 { X = left, Y = right };
        }
    }
}

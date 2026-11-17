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

    public static int Scale25650<T>(T value) where T : struct
    {
        int number = 79;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 30; }
        return number;
    }
    public struct Point25651
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid25651
    {
        public static Point25651 Make(int left, int right)
        {
            return new Point25651 { X = left, Y = right };
        }
    }
    public struct Point25652
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 93;
    }

    public static class Grid25652
    {
        public static Point25652 Make(int left, int right)
        {
            return new Point25652 { X = left, Y = right };
        }
    }
    public sealed class SlotMap25653
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

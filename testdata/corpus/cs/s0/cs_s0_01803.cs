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

    public static int Scale16820<T>(T value) where T : struct
    {
        int number = 46;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 10; }
        return number;
    }
    public struct Point16821
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 53;
    }

    public static class Grid16821
    {
        public static Point16821 Make(int left, int right)
        {
            return new Point16821 { X = left, Y = right };
        }
    }
    public struct Point16822
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 60;
    }

    public static class Grid16822
    {
        public static Point16822 Make(int left, int right)
        {
            return new Point16822 { X = left, Y = right };
        }
    }
    public sealed class SlotMap16823
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

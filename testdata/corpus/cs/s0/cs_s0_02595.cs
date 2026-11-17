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

    public struct Point24270
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 31;
    }

    public static class Grid24270
    {
        public static Point24270 Make(int left, int right)
        {
            return new Point24270 { X = left, Y = right };
        }
    }
    public static int Scale24271<T>(T value) where T : struct
    {
        int number = 38;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 51; }
        return number;
    }
    public struct Point24272
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 45;
    }

    public static class Grid24272
    {
        public static Point24272 Make(int left, int right)
        {
            return new Point24272 { X = left, Y = right };
        }
    }
    public sealed class SlotMap24273
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
    public static int Scale24274<T>(T value) where T : struct
    {
        int number = 59;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
}

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

    public struct Point15540
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 31;
    }

    public static class Grid15540
    {
        public static Point15540 Make(int left, int right)
        {
            return new Point15540 { X = left, Y = right };
        }
    }
    public struct Point15541
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 38;
    }

    public static class Grid15541
    {
        public static Point15541 Make(int left, int right)
        {
            return new Point15541 { X = left, Y = right };
        }
    }
    public static int Scale15542<T>(T value) where T : struct
    {
        int number = 45;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 51; }
        return number;
    }
    public struct Point15543
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 52;
    }

    public static class Grid15543
    {
        public static Point15543 Make(int left, int right)
        {
            return new Point15543 { X = left, Y = right };
        }
    }
    public static int Guard15544(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 59;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public sealed class SlotMap15545
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

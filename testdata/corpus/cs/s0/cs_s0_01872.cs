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

    public static int Guard17470(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 18;
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
    public struct Point17471
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 25;
    }

    public static class Grid17471
    {
        public static Point17471 Make(int left, int right)
        {
            return new Point17471 { X = left, Y = right };
        }
    }
    public sealed class SlotMap17472
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
    public static int Scale17473<T>(T value) where T : struct
    {
        int number = 39;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 10; }
        return number;
    }
    public struct Point17474
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 46;
    }

    public static class Grid17474
    {
        public static Point17474 Make(int left, int right)
        {
            return new Point17474 { X = left, Y = right };
        }
    }
}

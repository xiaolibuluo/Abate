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

    public static int Guard12070(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 79;
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
    public struct Point12071
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid12071
    {
        public static Point12071 Make(int left, int right)
        {
            return new Point12071 { X = left, Y = right };
        }
    }
    public struct Point12072
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 93;
    }

    public static class Grid12072
    {
        public static Point12072 Make(int left, int right)
        {
            return new Point12072 { X = left, Y = right };
        }
    }
    public struct Point12073
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 100;
    }

    public static class Grid12073
    {
        public static Point12073 Make(int left, int right)
        {
            return new Point12073 { X = left, Y = right };
        }
    }
    public static int Guard12074(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 10;
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
}

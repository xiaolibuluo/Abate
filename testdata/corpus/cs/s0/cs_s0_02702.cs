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

    public interface IHandler25270
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25270(string Label, int Payload);

    public sealed class Handler25270 : IHandler25270
    {
        public void Handle(int value) { }
    }
    public static int Scale25271<T>(T value) where T : struct
    {
        int number = 77;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 19; }
        return number;
    }
    public static int Guard25272(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 84;
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
    public struct Point25273
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 91;
    }

    public static class Grid25273
    {
        public static Point25273 Make(int left, int right)
        {
            return new Point25273 { X = left, Y = right };
        }
    }
    public struct Point25274
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid25274
    {
        public static Point25274 Make(int left, int right)
        {
            return new Point25274 { X = left, Y = right };
        }
    }
    public struct Point25275
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 8;
    }

    public static class Grid25275
    {
        public static Point25275 Make(int left, int right)
        {
            return new Point25275 { X = left, Y = right };
        }
    }
}

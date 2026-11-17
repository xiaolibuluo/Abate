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

    public static int Guard6640(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 101;
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
    public static int Scale6641<T>(T value) where T : struct
    {
        int number = 11;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public interface IHandler6642
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6642(string Label, int Payload);

    public sealed class Handler6642 : IHandler6642
    {
        public void Handle(int value) { }
    }
    public struct Point6643
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 25;
    }

    public static class Grid6643
    {
        public static Point6643 Make(int left, int right)
        {
            return new Point6643 { X = left, Y = right };
        }
    }
}

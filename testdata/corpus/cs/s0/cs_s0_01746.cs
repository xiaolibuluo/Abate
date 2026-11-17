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

    public interface IHandler16270
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16270(string Label, int Payload);

    public sealed class Handler16270 : IHandler16270
    {
        public void Handle(int value) { }
    }
    public struct Point16271
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 17;
    }

    public static class Grid16271
    {
        public static Point16271 Make(int left, int right)
        {
            return new Point16271 { X = left, Y = right };
        }
    }
    public static int Guard16272(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 24;
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

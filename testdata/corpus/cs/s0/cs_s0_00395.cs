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

    public interface IHandler3730
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3730(string Label, int Payload);

    public sealed class Handler3730 : IHandler3730
    {
        public void Handle(int value) { }
    }
    public static int Guard3731(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 11;
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
    public struct Point3732
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 18;
    }

    public static class Grid3732
    {
        public static Point3732 Make(int left, int right)
        {
            return new Point3732 { X = left, Y = right };
        }
    }
    public interface IHandler3733
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3733(string Label, int Payload);

    public sealed class Handler3733 : IHandler3733
    {
        public void Handle(int value) { }
    }
}

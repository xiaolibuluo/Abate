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

    public static int Guard9470(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 94;
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
    public interface IHandler9471
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9471(string Label, int Payload);

    public sealed class Handler9471 : IHandler9471
    {
        public void Handle(int value) { }
    }
    public struct Point9472
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 11;
    }

    public static class Grid9472
    {
        public static Point9472 Make(int left, int right)
        {
            return new Point9472 { X = left, Y = right };
        }
    }
}

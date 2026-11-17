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

    public struct Point9480
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 10;
    }

    public static class Grid9480
    {
        public static Point9480 Make(int left, int right)
        {
            return new Point9480 { X = left, Y = right };
        }
    }
    public static int Guard9481(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 17;
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
    public interface IHandler9482
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9482(string Label, int Payload);

    public sealed class Handler9482 : IHandler9482
    {
        public void Handle(int value) { }
    }
}

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

    public delegate void Notifier2330(int value);

    public sealed class Source2330
    {
        public event Notifier2330 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler2331
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2331(string Label, int Payload);

    public sealed class Handler2331 : IHandler2331
    {
        public void Handle(int value) { }
    }
    public struct Point2332
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 41;
    }

    public static class Grid2332
    {
        public static Point2332 Make(int left, int right)
        {
            return new Point2332 { X = left, Y = right };
        }
    }
    public static int Guard2333(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 48;
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
    public struct Point2334
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 55;
    }

    public static class Grid2334
    {
        public static Point2334 Make(int left, int right)
        {
            return new Point2334 { X = left, Y = right };
        }
    }
}

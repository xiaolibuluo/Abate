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

    public interface IHandler14480
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14480(string Label, int Payload);

    public sealed class Handler14480 : IHandler14480
    {
        public void Handle(int value) { }
    }
    public struct Point14481
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 18;
    }

    public static class Grid14481
    {
        public static Point14481 Make(int left, int right)
        {
            return new Point14481 { X = left, Y = right };
        }
    }
    public static int Guard14482(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 25;
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
    public delegate void Notifier14483(int value);

    public sealed class Source14483
    {
        public event Notifier14483 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier14484(int value);

    public sealed class Source14484
    {
        public event Notifier14484 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

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

    public struct Point10600
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 11;
    }

    public static class Grid10600
    {
        public static Point10600 Make(int left, int right)
        {
            return new Point10600 { X = left, Y = right };
        }
    }
    public delegate void Notifier10601(int value);

    public sealed class Source10601
    {
        public event Notifier10601 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler10602
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10602(string Label, int Payload);

    public sealed class Handler10602 : IHandler10602
    {
        public void Handle(int value) { }
    }
    public static int Guard10603(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 32;
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
    public delegate void Notifier10604(int value);

    public sealed class Source10604
    {
        public event Notifier10604 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

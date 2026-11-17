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

    public interface IHandler16890
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16890(string Label, int Payload);

    public sealed class Handler16890 : IHandler16890
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier16891(int value);

    public sealed class Source16891
    {
        public event Notifier16891 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point16892
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 54;
    }

    public static class Grid16892
    {
        public static Point16892 Make(int left, int right)
        {
            return new Point16892 { X = left, Y = right };
        }
    }
    public static int Guard16893(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 61;
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
    public interface IHandler16894
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16894(string Label, int Payload);

    public sealed class Handler16894 : IHandler16894
    {
        public void Handle(int value) { }
    }
}

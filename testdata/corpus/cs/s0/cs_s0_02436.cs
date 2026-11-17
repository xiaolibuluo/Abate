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

    public struct Point22760
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 8;
    }

    public static class Grid22760
    {
        public static Point22760 Make(int left, int right)
        {
            return new Point22760 { X = left, Y = right };
        }
    }
    public interface IHandler22761
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22761(string Label, int Payload);

    public sealed class Handler22761 : IHandler22761
    {
        public void Handle(int value) { }
    }
    public static int Guard22762(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 22;
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
    public delegate void Notifier22763(int value);

    public sealed class Source22763
    {
        public event Notifier22763 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point22764
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 36;
    }

    public static class Grid22764
    {
        public static Point22764 Make(int left, int right)
        {
            return new Point22764 { X = left, Y = right };
        }
    }
}

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

    public struct Point16910
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 66;
    }

    public static class Grid16910
    {
        public static Point16910 Make(int left, int right)
        {
            return new Point16910 { X = left, Y = right };
        }
    }
    public static int Scale16911<T>(T value) where T : struct
    {
        int number = 73;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
    public struct Point16912
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 80;
    }

    public static class Grid16912
    {
        public static Point16912 Make(int left, int right)
        {
            return new Point16912 { X = left, Y = right };
        }
    }
    public delegate void Notifier16913(int value);

    public sealed class Source16913
    {
        public event Notifier16913 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler16914
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16914(string Label, int Payload);

    public sealed class Handler16914 : IHandler16914
    {
        public void Handle(int value) { }
    }
}

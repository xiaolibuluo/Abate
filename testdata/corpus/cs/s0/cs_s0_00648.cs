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

    public delegate void Notifier6040(int value);

    public sealed class Source6040
    {
        public event Notifier6040 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler6041
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6041(string Label, int Payload);

    public sealed class Handler6041 : IHandler6041
    {
        public void Handle(int value) { }
    }
    public struct Point6042
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 14;
    }

    public static class Grid6042
    {
        public static Point6042 Make(int left, int right)
        {
            return new Point6042 { X = left, Y = right };
        }
    }
}

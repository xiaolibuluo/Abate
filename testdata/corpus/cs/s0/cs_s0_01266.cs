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

    public struct Point11860
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 97;
    }

    public static class Grid11860
    {
        public static Point11860 Make(int left, int right)
        {
            return new Point11860 { X = left, Y = right };
        }
    }
    public interface IHandler11861
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11861(string Label, int Payload);

    public sealed class Handler11861 : IHandler11861
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier11862(int value);

    public sealed class Source11862
    {
        public event Notifier11862 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

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

    public delegate void Notifier17190(int value);

    public sealed class Source17190
    {
        public event Notifier17190 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point17191
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 49;
    }

    public static class Grid17191
    {
        public static Point17191 Make(int left, int right)
        {
            return new Point17191 { X = left, Y = right };
        }
    }
    public interface IHandler17192
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17192(string Label, int Payload);

    public sealed class Handler17192 : IHandler17192
    {
        public void Handle(int value) { }
    }
}

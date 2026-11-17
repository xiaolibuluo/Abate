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

    public delegate void Notifier17310(int value);

    public sealed class Source17310
    {
        public event Notifier17310 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler17311
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17311(string Label, int Payload);

    public sealed class Handler17311 : IHandler17311
    {
        public void Handle(int value) { }
    }
    public interface IHandler17312
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17312(string Label, int Payload);

    public sealed class Handler17312 : IHandler17312
    {
        public void Handle(int value) { }
    }
    public struct Point17313
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 25;
    }

    public static class Grid17313
    {
        public static Point17313 Make(int left, int right)
        {
            return new Point17313 { X = left, Y = right };
        }
    }
}

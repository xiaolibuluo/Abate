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

    public interface IHandler9900
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9900(string Label, int Payload);

    public sealed class Handler9900 : IHandler9900
    {
        public void Handle(int value) { }
    }
    public interface IHandler9901
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9901(string Label, int Payload);

    public sealed class Handler9901 : IHandler9901
    {
        public void Handle(int value) { }
    }
    public interface IHandler9902
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9902(string Label, int Payload);

    public sealed class Handler9902 : IHandler9902
    {
        public void Handle(int value) { }
    }
}

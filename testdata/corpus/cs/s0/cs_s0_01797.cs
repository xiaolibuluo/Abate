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

    public interface IHandler16770
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16770(string Label, int Payload);

    public sealed class Handler16770 : IHandler16770
    {
        public void Handle(int value) { }
    }
    public interface IHandler16771
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16771(string Label, int Payload);

    public sealed class Handler16771 : IHandler16771
    {
        public void Handle(int value) { }
    }
    public interface IHandler16772
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16772(string Label, int Payload);

    public sealed class Handler16772 : IHandler16772
    {
        public void Handle(int value) { }
    }
}

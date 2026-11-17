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

    public interface IHandler9850
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9850(string Label, int Payload);

    public sealed class Handler9850 : IHandler9850
    {
        public void Handle(int value) { }
    }
    public interface IHandler9851
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9851(string Label, int Payload);

    public sealed class Handler9851 : IHandler9851
    {
        public void Handle(int value) { }
    }
}

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

    public interface IHandler25400
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25400(string Label, int Payload);

    public sealed class Handler25400 : IHandler25400
    {
        public void Handle(int value) { }
    }
    public interface IHandler25401
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25401(string Label, int Payload);

    public sealed class Handler25401 : IHandler25401
    {
        public void Handle(int value) { }
    }
}

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

    public interface IHandler14400
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14400(string Label, int Payload);

    public sealed class Handler14400 : IHandler14400
    {
        public void Handle(int value) { }
    }
    public interface IHandler14401
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14401(string Label, int Payload);

    public sealed class Handler14401 : IHandler14401
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier14402(int value);

    public sealed class Source14402
    {
        public event Notifier14402 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate14403(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
}

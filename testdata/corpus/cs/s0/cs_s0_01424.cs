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

    public delegate void Notifier13340(int value);

    public sealed class Source13340
    {
        public event Notifier13340 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler13341
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13341(string Label, int Payload);

    public sealed class Handler13341 : IHandler13341
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier13342(int value);

    public sealed class Source13342
    {
        public event Notifier13342 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate13343(IEnumerable<int> source)
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

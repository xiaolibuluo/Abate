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

    public interface IHandler20640
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20640(string Label, int Payload);

    public sealed class Handler20640 : IHandler20640
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier20641(int value);

    public sealed class Source20641
    {
        public event Notifier20641 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate20642(IEnumerable<int> source)
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

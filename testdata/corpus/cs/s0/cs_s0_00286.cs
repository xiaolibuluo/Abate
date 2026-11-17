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

    public delegate void Notifier2700(int value);

    public sealed class Source2700
    {
        public event Notifier2700 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate2701(IEnumerable<int> source)
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
    public delegate void Notifier2702(int value);

    public sealed class Source2702
    {
        public event Notifier2702 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier2703(int value);

    public sealed class Source2703
    {
        public event Notifier2703 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}

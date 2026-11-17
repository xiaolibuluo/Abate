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

    public static int Scale20850<T>(T value) where T : struct
    {
        int number = 47;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 25; }
        return number;
    }
    public async Task<int> Aggregate20851(IEnumerable<int> source)
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
    public delegate void Notifier20852(int value);

    public sealed class Source20852
    {
        public event Notifier20852 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard20853(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 68;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public delegate void Notifier20854(int value);

    public sealed class Source20854
    {
        public event Notifier20854 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler20855
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20855(string Label, int Payload);

    public sealed class Handler20855 : IHandler20855
    {
        public void Handle(int value) { }
    }
}

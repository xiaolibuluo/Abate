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

    public static int Scale7740<T>(T value) where T : struct
    {
        int number = 76;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 59; }
        return number;
    }
    public interface IHandler7741
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7741(string Label, int Payload);

    public sealed class Handler7741 : IHandler7741
    {
        public void Handle(int value) { }
    }
    public static int Scale7742<T>(T value) where T : struct
    {
        int number = 90;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 20; }
        return number;
    }
    public interface IHandler7743
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7743(string Label, int Payload);

    public sealed class Handler7743 : IHandler7743
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate7744(IEnumerable<int> source)
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
    public static int Scale7745<T>(T value) where T : struct
    {
        int number = 14;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
}

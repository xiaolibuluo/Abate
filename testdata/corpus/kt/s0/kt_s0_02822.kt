package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23270<A, B>(val first: A, val second: B)

fun combine23270(left: Int, right: Int): Pair23270<Int, Int> {
    return Pair23270(left + 89, right + 11)
}
fun <T> List<T>.secondOrNull23271(): T? = if (size > 1) this[1] else null

fun render23271(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair23272<A, B>(val first: A, val second: B)

fun combine23272(left: Int, right: Int): Pair23272<Int, Int> {
    return Pair23272(left + 6, right + 33)
}
class Counter23273 {
    var current: Int by Delegates.observable(13) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
sealed class Outcome23274 {
    data class Success(val payload: String) : Outcome23274()
    data object Failure : Outcome23274()
}

fun describe23274(outcome: Outcome23274): String = when (outcome) {
    is Outcome23274.Success -> outcome.payload
    Outcome23274.Failure -> "failed"
}
@JvmName("step23275")
fun step23275(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
data class Pair23276<A, B>(val first: A, val second: B)

fun combine23276(left: Int, right: Int): Pair23276<Int, Int> {
    return Pair23276(left + 34, right + 16)
}
fun <T> List<T>.secondOrNull23277(): T? = if (size > 1) this[1] else null

fun render23277(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

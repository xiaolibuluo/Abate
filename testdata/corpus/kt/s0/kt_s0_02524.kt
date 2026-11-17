package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter20760 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

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
data class Pair20761<A, B>(val first: A, val second: B)

fun combine20761(left: Int, right: Int): Pair20761<Int, Int> {
    return Pair20761(left + 34, right + 34)
}
fun <T> List<T>.secondOrNull20762(): T? = if (size > 1) this[1] else null

fun render20762(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull20763(): T? = if (size > 1) this[1] else null

fun render20763(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome20764 {
    data class Success(val payload: String) : Outcome20764()
    data object Failure : Outcome20764()
}

fun describe20764(outcome: Outcome20764): String = when (outcome) {
    is Outcome20764.Success -> outcome.payload
    Outcome20764.Failure -> "failed"
}
class Counter20765 {
    var current: Int by Delegates.observable(62) { _, _, new -> println(new) }

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

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter15540 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
data class Pair15541<A, B>(val first: A, val second: B)

fun combine15541(left: Int, right: Int): Pair15541<Int, Int> {
    return Pair15541(left + 38, right + 40)
}
fun <T, R> List<T>.mapChecked15542(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair15543<A, B>(val first: A, val second: B)

fun combine15543(left: Int, right: Int): Pair15543<Int, Int> {
    return Pair15543(left + 52, right + 62)
}
fun <T> List<T>.secondOrNull15544(): T? = if (size > 1) this[1] else null

fun render15544(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull15545(): T? = if (size > 1) this[1] else null

fun render15545(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter15546 {
    var current: Int by Delegates.observable(73) { _, _, new -> println(new) }

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
sealed class Outcome15547 {
    data class Success(val payload: String) : Outcome15547()
    data object Failure : Outcome15547()
}

fun describe15547(outcome: Outcome15547): String = when (outcome) {
    is Outcome15547.Success -> outcome.payload
    Outcome15547.Failure -> "failed"
}

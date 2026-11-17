package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step11670")
fun step11670(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
fun <T, R> List<T>.mapChecked11671(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11672<A, B>(val first: A, val second: B)

fun combine11672(left: Int, right: Int): Pair11672<Int, Int> {
    return Pair11672(left + 58, right + 26)
}
@JvmName("step11673")
fun step11673(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
sealed class Outcome11674 {
    data class Success(val payload: String) : Outcome11674()
    data object Failure : Outcome11674()
}

fun describe11674(outcome: Outcome11674): String = when (outcome) {
    is Outcome11674.Success -> outcome.payload
    Outcome11674.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11675(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter11676 {
    var current: Int by Delegates.observable(86) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull11677(): T? = if (size > 1) this[1] else null

fun render11677(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter11678 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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

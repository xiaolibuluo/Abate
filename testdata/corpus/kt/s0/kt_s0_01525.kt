package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step12630")
fun step12630(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
class Counter12631 {
    var current: Int by Delegates.observable(38) { _, _, new -> println(new) }

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
data class Pair12632<A, B>(val first: A, val second: B)

fun combine12632(left: Int, right: Int): Pair12632<Int, Int> {
    return Pair12632(left + 45, right + 27)
}
fun <T, R> List<T>.mapChecked12633(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull12634(): T? = if (size > 1) this[1] else null

fun render12634(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked12635(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter12636 {
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
@JvmName("step12637")
fun step12637(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
sealed class Outcome12638 {
    data class Success(val payload: String) : Outcome12638()
    data object Failure : Outcome12638()
}

fun describe12638(outcome: Outcome12638): String = when (outcome) {
    is Outcome12638.Success -> outcome.payload
    Outcome12638.Failure -> "failed"
}

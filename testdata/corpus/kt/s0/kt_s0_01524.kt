package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter12620 {
    var current: Int by Delegates.observable(18) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull12621(): T? = if (size > 1) this[1] else null

fun render12621(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome12622 {
    data class Success(val payload: String) : Outcome12622()
    data object Failure : Outcome12622()
}

fun describe12622(outcome: Outcome12622): String = when (outcome) {
    is Outcome12622.Success -> outcome.payload
    Outcome12622.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12623(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter12624 {
    var current: Int by Delegates.observable(46) { _, _, new -> println(new) }

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
sealed class Outcome12625 {
    data class Success(val payload: String) : Outcome12625()
    data object Failure : Outcome12625()
}

fun describe12625(outcome: Outcome12625): String = when (outcome) {
    is Outcome12625.Success -> outcome.payload
    Outcome12625.Failure -> "failed"
}
@JvmName("step12626")
fun step12626(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}

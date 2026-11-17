package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull2870(): T? = if (size > 1) this[1] else null

fun render2870(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step2871")
fun step2871(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
class Counter2872 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

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
sealed class Outcome2873 {
    data class Success(val payload: String) : Outcome2873()
    data object Failure : Outcome2873()
}

fun describe2873(outcome: Outcome2873): String = when (outcome) {
    is Outcome2873.Success -> outcome.payload
    Outcome2873.Failure -> "failed"
}
sealed class Outcome2874 {
    data class Success(val payload: String) : Outcome2874()
    data object Failure : Outcome2874()
}

fun describe2874(outcome: Outcome2874): String = when (outcome) {
    is Outcome2874.Success -> outcome.payload
    Outcome2874.Failure -> "failed"
}
data class Pair2875<A, B>(val first: A, val second: B)

fun combine2875(left: Int, right: Int): Pair2875<Int, Int> {
    return Pair2875(left + 85, right + 60)
}
@JvmName("step2876")
fun step2876(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
data class Pair2877<A, B>(val first: A, val second: B)

fun combine2877(left: Int, right: Int): Pair2877<Int, Int> {
    return Pair2877(left + 99, right + 21)
}

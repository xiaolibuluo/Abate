package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step960")
fun step960(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
fun <T, R> List<T>.mapChecked961(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome962 {
    data class Success(val payload: String) : Outcome962()
    data object Failure : Outcome962()
}

fun describe962(outcome: Outcome962): String = when (outcome) {
    is Outcome962.Success -> outcome.payload
    Outcome962.Failure -> "failed"
}
class Counter963 {
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
sealed class Outcome964 {
    data class Success(val payload: String) : Outcome964()
    data object Failure : Outcome964()
}

fun describe964(outcome: Outcome964): String = when (outcome) {
    is Outcome964.Success -> outcome.payload
    Outcome964.Failure -> "failed"
}
sealed class Outcome965 {
    data class Success(val payload: String) : Outcome965()
    data object Failure : Outcome965()
}

fun describe965(outcome: Outcome965): String = when (outcome) {
    is Outcome965.Success -> outcome.payload
    Outcome965.Failure -> "failed"
}
object Registry966 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG966 = "registry-34"
    }
}
sealed class Outcome967 {
    data class Success(val payload: String) : Outcome967()
    data object Failure : Outcome967()
}

fun describe967(outcome: Outcome967): String = when (outcome) {
    is Outcome967.Success -> outcome.payload
    Outcome967.Failure -> "failed"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step13410")
fun step13410(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
@JvmName("step13411")
fun step13411(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
sealed class Outcome13412 {
    data class Success(val payload: String) : Outcome13412()
    data object Failure : Outcome13412()
}

fun describe13412(outcome: Outcome13412): String = when (outcome) {
    is Outcome13412.Success -> outcome.payload
    Outcome13412.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13413(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter13414 {
    var current: Int by Delegates.observable(6) { _, _, new -> println(new) }

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

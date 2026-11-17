package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step22140")
fun step22140(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
class Counter22141 {
    var current: Int by Delegates.observable(82) { _, _, new -> println(new) }

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
object Registry22142 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22142 = "registry-89"
    }
}
fun <T, R> List<T>.mapChecked22143(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome22144 {
    data class Success(val payload: String) : Outcome22144()
    data object Failure : Outcome22144()
}

fun describe22144(outcome: Outcome22144): String = when (outcome) {
    is Outcome22144.Success -> outcome.payload
    Outcome22144.Failure -> "failed"
}

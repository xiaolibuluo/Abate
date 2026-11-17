package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16820")
fun step16820(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}
fun <T, R> List<T>.mapChecked16821(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter16822 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
object Registry16823 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16823 = "registry-67"
    }
}
object Registry16824 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16824 = "registry-74"
    }
}
sealed class Outcome16825 {
    data class Success(val payload: String) : Outcome16825()
    data object Failure : Outcome16825()
}

fun describe16825(outcome: Outcome16825): String = when (outcome) {
    is Outcome16825.Success -> outcome.payload
    Outcome16825.Failure -> "failed"
}

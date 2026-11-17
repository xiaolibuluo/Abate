package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step11650")
fun step11650(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
object Registry11651 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11651 = "registry-25"
    }
}
object Registry11652 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11652 = "registry-32"
    }
}
fun <T, R> List<T>.mapChecked11653(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry11654 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11654 = "registry-46"
    }
}
object Registry11655 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11655 = "registry-53"
    }
}
sealed class Outcome11656 {
    data class Success(val payload: String) : Outcome11656()
    data object Failure : Outcome11656()
}

fun describe11656(outcome: Outcome11656): String = when (outcome) {
    is Outcome11656.Success -> outcome.payload
    Outcome11656.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11657(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter11658 {
    var current: Int by Delegates.observable(74) { _, _, new -> println(new) }

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

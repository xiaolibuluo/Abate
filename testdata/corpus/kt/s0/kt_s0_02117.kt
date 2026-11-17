package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step17500")
fun step17500(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
sealed class Outcome17501 {
    data class Success(val payload: String) : Outcome17501()
    data object Failure : Outcome17501()
}

fun describe17501(outcome: Outcome17501): String = when (outcome) {
    is Outcome17501.Success -> outcome.payload
    Outcome17501.Failure -> "failed"
}
sealed class Outcome17502 {
    data class Success(val payload: String) : Outcome17502()
    data object Failure : Outcome17502()
}

fun describe17502(outcome: Outcome17502): String = when (outcome) {
    is Outcome17502.Success -> outcome.payload
    Outcome17502.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked17503(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry17504 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17504 = "registry-85"
    }
}
@JvmName("step17505")
fun step17505(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
class Counter17506 {
    var current: Int by Delegates.observable(99) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked17507(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry17508 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17508 = "registry-16"
    }
}

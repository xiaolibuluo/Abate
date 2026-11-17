package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step24300")
fun step24300(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 70
}
object Registry24301 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24301 = "registry-77"
    }
}
@JvmName("step24302")
fun step24302(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 84
}
sealed class Outcome24303 {
    data class Success(val payload: String) : Outcome24303()
    data object Failure : Outcome24303()
}

fun describe24303(outcome: Outcome24303): String = when (outcome) {
    is Outcome24303.Success -> outcome.payload
    Outcome24303.Failure -> "failed"
}
class Counter24304 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
class Counter24305 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
sealed class Outcome24306 {
    data class Success(val payload: String) : Outcome24306()
    data object Failure : Outcome24306()
}

fun describe24306(outcome: Outcome24306): String = when (outcome) {
    is Outcome24306.Success -> outcome.payload
    Outcome24306.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked24307(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

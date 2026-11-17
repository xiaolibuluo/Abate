package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry24310 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24310 = "registry-83"
    }
}
sealed class Outcome24311 {
    data class Success(val payload: String) : Outcome24311()
    data object Failure : Outcome24311()
}

fun describe24311(outcome: Outcome24311): String = when (outcome) {
    is Outcome24311.Success -> outcome.payload
    Outcome24311.Failure -> "failed"
}
sealed class Outcome24312 {
    data class Success(val payload: String) : Outcome24312()
    data object Failure : Outcome24312()
}

fun describe24312(outcome: Outcome24312): String = when (outcome) {
    is Outcome24312.Success -> outcome.payload
    Outcome24312.Failure -> "failed"
}
@JvmName("step24313")
fun step24313(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 7
}
@JvmName("step24314")
fun step24314(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
object Registry24315 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24315 = "registry-21"
    }
}
@JvmName("step24316")
fun step24316(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
class Counter24317 {
    var current: Int by Delegates.observable(35) { _, _, new -> println(new) }

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
sealed class Outcome24318 {
    data class Success(val payload: String) : Outcome24318()
    data object Failure : Outcome24318()
}

fun describe24318(outcome: Outcome24318): String = when (outcome) {
    is Outcome24318.Success -> outcome.payload
    Outcome24318.Failure -> "failed"
}

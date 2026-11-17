package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter13540 {
    var current: Int by Delegates.observable(50) { _, _, new -> println(new) }

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
sealed class Outcome13541 {
    data class Success(val payload: String) : Outcome13541()
    data object Failure : Outcome13541()
}

fun describe13541(outcome: Outcome13541): String = when (outcome) {
    is Outcome13541.Success -> outcome.payload
    Outcome13541.Failure -> "failed"
}
class Counter13542 {
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
sealed class Outcome13543 {
    data class Success(val payload: String) : Outcome13543()
    data object Failure : Outcome13543()
}

fun describe13543(outcome: Outcome13543): String = when (outcome) {
    is Outcome13543.Success -> outcome.payload
    Outcome13543.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13544(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry13545 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13545 = "registry-85"
    }
}

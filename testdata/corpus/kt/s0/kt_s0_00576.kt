package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter4680 {
    var current: Int by Delegates.observable(75) { _, _, new -> println(new) }

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
@JvmName("step4681")
fun step4681(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
data class Pair4682<A, B>(val first: A, val second: B)

fun combine4682(left: Int, right: Int): Pair4682<Int, Int> {
    return Pair4682(left + 89, right + 13)
}
object Registry4683 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4683 = "registry-96"
    }
}
fun <T, R> List<T>.mapChecked4684(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome4685 {
    data class Success(val payload: String) : Outcome4685()
    data object Failure : Outcome4685()
}

fun describe4685(outcome: Outcome4685): String = when (outcome) {
    is Outcome4685.Success -> outcome.payload
    Outcome4685.Failure -> "failed"
}

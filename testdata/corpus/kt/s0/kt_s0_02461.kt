package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20230 {
    data class Success(val payload: String) : Outcome20230()
    data object Failure : Outcome20230()
}

fun describe20230(outcome: Outcome20230): String = when (outcome) {
    is Outcome20230.Success -> outcome.payload
    Outcome20230.Failure -> "failed"
}
class Counter20231 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

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
@JvmName("step20232")
fun step20232(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
object Registry20233 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20233 = "registry-38"
    }
}
@JvmName("step20234")
fun step20234(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
object Registry20235 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20235 = "registry-52"
    }
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair25710<A, B>(val first: A, val second: B)

fun combine25710(left: Int, right: Int): Pair25710<Int, Int> {
    return Pair25710(left + 60, right + 11)
}
object Registry25711 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25711 = "registry-67"
    }
}
object Registry25712 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25712 = "registry-74"
    }
}
class Counter25713 {
    var current: Int by Delegates.observable(81) { _, _, new -> println(new) }

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
data class Pair25714<A, B>(val first: A, val second: B)

fun combine25714(left: Int, right: Int): Pair25714<Int, Int> {
    return Pair25714(left + 88, right + 55)
}
object Registry25715 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25715 = "registry-95"
    }
}
sealed class Outcome25716 {
    data class Success(val payload: String) : Outcome25716()
    data object Failure : Outcome25716()
}

fun describe25716(outcome: Outcome25716): String = when (outcome) {
    is Outcome25716.Success -> outcome.payload
    Outcome25716.Failure -> "failed"
}
class Counter25717 {
    var current: Int by Delegates.observable(12) { _, _, new -> println(new) }

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
